#import "window.h"
#import <CoreGraphics/CoreGraphics.h>
#import <Foundation/Foundation.h>

int GetActiveWindowDetails(char** appNameOut, char** windowTitleOut) {
    @autoreleasepool {
        CFArrayRef windowList = CGWindowListCopyWindowInfo(kCGWindowListOptionOnScreenOnly | kCGWindowListExcludeDesktopElements, kCGNullWindowID);
        NSArray* windows = (__bridge NSArray*)windowList;
        
        int success = 0;

        for (NSDictionary* window in windows) {
            NSInteger layer = [[window objectForKey:(id)kCGWindowLayer] integerValue];
            if (layer == 0) { 
                NSString* ownerName = [window objectForKey:(id)kCGWindowOwnerName];
                NSString* windowName = [window objectForKey:(id)kCGWindowName];
                
                if (ownerName) {
                    *appNameOut = strdup([ownerName UTF8String]);
                    
                    if (windowName && windowName.length > 0) {
                        *windowTitleOut = strdup([windowName UTF8String]);
                    } else {
                        *windowTitleOut = strdup(""); 
                    }
                    success = 1;
                    break;
                }
            }
        }
        
        CFRelease(windowList);
        return success;
    }
}
