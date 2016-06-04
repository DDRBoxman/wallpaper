package wallpaper

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>

void
setWallpaper(char* name) {
     NSString *nsName =  [NSString stringWithUTF8String:name];
     NSString *urlString = [NSString stringWithFormat:@"file://%@", nsName];
     NSURL *imageURL = [NSURL URLWithString:urlString];
    [[NSWorkspace sharedWorkspace] setDesktopImageURL:imageURL
                                   forScreen:[NSScreen mainScreen]
                                   options:nil
                                   error:nil];
    free(name);
}
*/
import "C"

func Set(name string) {
	C.setWallpaper(C.CString(name))
}
