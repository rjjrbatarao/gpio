package gpio

import (
	"syscall"
)

func mySelect(nfd int, readfds *syscall.FdSet, writefds *syscall.FdSet, exceptfds *syscall.FdSet, timeout *syscall.Timeval) (n int, err error) {
    for {
        n, err := syscall.Select(nfd, readfds, writefds, exceptfds, timeout)
        if err == syscall.EINTR {
            // The system call was interrupted by a signal.
            // Loop and try again.
            continue
        }
        if err != nil {
            // A real error occurred.
            return n, err
        }
        // Success (or timeout reached without an error other than EINTR).
        return n, nil
    }
}

func doSelect(nfd int, r *syscall.FdSet, w *syscall.FdSet, e *syscall.FdSet, timeout *syscall.Timeval) (changed bool, err error) {
	n, err := mySelect(nfd, r, w, e, timeout)
	if err != nil {
		return false, err
	}
	if n != 0 {
		return true, nil
	}
	return false, nil
}
