>>>>>>> interrupt

http://stackoverflow.com/questions/24466375/disabling-interrupts-for-synchornization-in-kernel-code
http://stackoverflow.com/questions/23754995/linux-thread-sleep-vs-read?answertab=active#tab-top


This necessary because the scheduler can be entered via both software and hardware interrupts.
Software interrupts, eg. sleep() calls and inter-thread comms calls, (eg semaphore, condvar or event 
signal), may change the set of running threads and so will request a scheduler run. 
These calls have thread/process context and happen whenever they call into the kernel.

Hardware interrupts, eg. KB, mouse, disk, NIC cause drivers to run and the driver may well
 wish to change the set of running threads by running the scheduler, eg. to make a thread 
ready that was blocked waiting for a disk read. Hardware interrupts have no thread/process 
context and can happen at any time while interrupts are enabled.

There are sections of scheduler code/data that are not reentrant. 
If interrupts are not briefly disabled for those sections, 
chaos will ensue when the scheduler is interrupted by hardware and reentered from a driver.