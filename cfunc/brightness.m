#import "brightness.h"
#import <CoreGraphics/CoreGraphics.h>

// Declare both private macOS functions
extern int DisplayServicesSetBrightness(CGDirectDisplayID display, float brightness);
extern int DisplayServicesGetBrightness(CGDirectDisplayID display, float *brightness);

void SetBrightness(float level) {
    if (level < 0.0f) level = 0.0f;
    if (level > 1.0f) level = 1.0f;

    CGDirectDisplayID mainDisplay = CGMainDisplayID();
    DisplayServicesSetBrightness(mainDisplay, level);
}

float GetBrightness(void) {
    float currentLevel = 1.0f;
    CGDirectDisplayID mainDisplay = CGMainDisplayID();
    
    // Pass the pointer to currentLevel so Apple's API can fill it
    DisplayServicesGetBrightness(mainDisplay, &currentLevel);
    return currentLevel;
}
