package wallpaper

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>

void
setWallpaper(void) {
     NSURL *imageURL = [NSURL URLWithString:@"file:///Users/ddrboxman/Pictures/0491C000395F7DB9E4EC06334FD01DC5.png"];
    [[NSWorkspace sharedWorkspace] setDesktopImageURL:imageURL
                                   forScreen:[NSScreen mainScreen]
                                   options:nil
                                   error:nil];

}
*/
import "C"

func Set() {
	C.setWallpaper()
}
