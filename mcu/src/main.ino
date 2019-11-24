#include "Arduino.h"
#include <Adafruit_NeoPixel.h>

#define PIN D1
#define LED_COUNT 106

// size for each serial payload
#define PAYLOAD_SIZE 8

Adafruit_NeoPixel strip = Adafruit_NeoPixel(LED_COUNT, PIN, NEO_GRB + NEO_KHZ800);

// init team colors to black (off)
uint32_t teamAPrimaryColor = strip.Color(0, 0, 0);
uint32_t teamASecondaryColor = strip.Color(0, 0, 0);

uint32_t teamBPrimaryColor = strip.Color(0, 0, 0);
uint32_t teamBSecondaryColor = strip.Color(0, 0, 0);

void setup() {
  Serial.setTimeout(100);
  Serial.begin(115200);

  strip.begin();
  strip.show(); // initialize all pixels to 'off'
}

void loop() {
  if (Serial.available() >= PAYLOAD_SIZE) {
    uint8_t buf[PAYLOAD_SIZE] = {0};
    Serial.readBytes(buf, PAYLOAD_SIZE);

    // the first element is the command
    switch (buf[0]) {
      case 1:
        // set team scored with fanfare
        showFanfareForTeam(buf[1]);
        break;
      case 2:
        // update scores
        updateScoreStates(buf[1], buf[2]);
        break;
      case 3:
        // set that a team has won
        setWinner(buf[1]);
        break;
      case 4:
        // set colors for a team
        setColorsForTeam(buf[1], buf[2], buf[3], buf[4], buf[5], buf[6], buf[7]);
        break;
      case 5:
        // for color tuning
        setColor(buf[1], buf[2], buf[3]);
        break;
    }
  }
}

// for tuning
void setColor(uint8_t r, uint8_t g, uint8_t b) {
  colorWipe(strip.Color(r, g, b), 50);
}

void showFanfareForTeam(uint8_t teamID) {
  uint32_t primary;
  uint32_t secondary;

  switch (teamID) {
    case 1:
      Serial.println("Showing fanfare for Team A");
      primary = teamAPrimaryColor;
      secondary = teamASecondaryColor;
      break;
    case 2:
      Serial.println("Showing fanfare for Team B");
      primary = teamBPrimaryColor;
      secondary = teamBSecondaryColor;
      break;
    default:
      return;
  }

  colorWipe(primary, 20);
  colorWipe(secondary, 20);

  for(int i=0; i<12; i++) {
      theaterChase(primary, 50);
      theaterChase(secondary, 50);
  }

  colorWipe(primary, 20);
  colorWipe(secondary, 20);
}

void updateScoreStates(uint8_t teamA, uint8_t teamB) {
  Serial.println("Updating scores");

  if (teamA == 0 && teamB == 0) {
    // prevent divide by zero, set them equal
    teamA = 1;
    teamB = 1;
  }

  float aRatio = float(teamA) / (teamA + teamB);
  // there are 16 lights on the ring under the star
  // the team that's winning should get the star
  // if they're tied, we split it
  bool splitStar = teamA == teamB;
  int nonStarPixels = strip.numPixels() - 16;

  // figure out the distributions
  int aTeamPixels = nonStarPixels * aRatio;

  // distribute the colors for team A
  for(int i=0; i<aTeamPixels; i++) {
    strip.setPixelColor(i, teamAPrimaryColor);
  }

  // distribute the colors for team B
  // by filling in the rest
  // this also keeps us from having to round on our float
  // to int conversions
  for(int i=aTeamPixels; i<nonStarPixels; i++) {
    strip.setPixelColor(i, teamBPrimaryColor);
  }

  if (splitStar) {
    // split the star colors
    for(int i=0; i<8; i++) {
      strip.setPixelColor(i + nonStarPixels, teamAPrimaryColor);
    }

    for(int i=8; i<16; i++) {
      strip.setPixelColor(i + nonStarPixels, teamBPrimaryColor);
    }
  } else {
    // color the star the leader's primary color
    uint32_t teamLeadingPrimary;

    if (teamA > teamB) {
      teamLeadingPrimary = teamAPrimaryColor;
    } else {
      teamLeadingPrimary = teamBPrimaryColor;
    }

    for(int i=0; i<16; i++) {
      strip.setPixelColor(i + nonStarPixels, teamLeadingPrimary);
    }
  }

  strip.show();
}

void setWinner(uint8_t teamID) {
    switch (teamID) {
    case 1:
      Serial.println("Setting that Team A has won");
      showFanfareForTeam(1);
      updateScoreStates(1, 0);
      break;
    case 2:
      Serial.println("Setting that Team B has won");
      showFanfareForTeam(2);
      updateScoreStates(0, 1);
      break;
  }
}

void setColorsForTeam(uint8_t teamID, uint8_t pr, uint8_t pg, uint8_t pb, uint8_t sr, uint8_t sg, uint8_t sb) {
  // tweak the colors by using the gamma table
  uint32_t primary = strip.gamma32(strip.Color(pr, pg, pb));
  uint32_t secondary = strip.gamma32(strip.Color(sr, sg, sb));

  switch (teamID) {
    case 1:
      Serial.println("Setting team A colors");
      teamAPrimaryColor = primary;
      teamASecondaryColor = secondary;
      break;
    case 2:
      Serial.println("Setting team B colors");
      teamBPrimaryColor = primary;
      teamBSecondaryColor = secondary;
      break;
  }
}

// Fill the dots one after the other with a color
void colorWipe(uint32_t c, uint8_t wait) {
  for(uint16_t i=0; i<strip.numPixels(); i++) {
    strip.setPixelColor(i, c);
    strip.show();
    delay(wait);
  }
}

//Theatre-style crawling lights.
void theaterChase(uint32_t c, uint8_t wait) {
  for (int j=0; j<10; j++) {  //do 10 cycles of chasing
    for (int q=0; q < 3; q++) {
      for (uint16_t i=0; i < strip.numPixels(); i=i+3) {
        strip.setPixelColor(i+q, c);    //turn every third pixel on
      }
      strip.show();

      delay(wait);

      for (uint16_t i=0; i < strip.numPixels(); i=i+3) {
        strip.setPixelColor(i+q, 0);        //turn every third pixel off
      }
    }
  }
}
