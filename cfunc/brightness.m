#include "brightness.h"
#include <CoreGraphics/CoreGraphics.h>

// Declare the private macOS function
extern int DisplayServicesSetBrightness(CGDirectDisplayID display, float brightness);

void SetBrightness(float level) {
    if (level < 0.0f) level = 0.0f;
    if (level > 1.0f) level = 1.0f;

    // Grab the primary active display ID (e.g., your MacBook screen)
    CGDirectDisplayID mainDisplay = CGMainDisplayID();
    
    // Set the brightness
    DisplayServicesSetBrightness(mainDisplay, level);
}
