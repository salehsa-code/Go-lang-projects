>Task 1 - Matching Behaviour


**What happens if you remove the go-command from the Seek call in the main function?**

Nothing changes in the output since it's buffered, it only gurantees the ordering of names. Furthermore, Seek will be run one at a time and waited for, since it won't be called in the goroutine 

**What happens if you switch the declaration wg := new(sync.WaitGroup)** 
**to var wg sync.WaitGroup and the parameter wg *sync.WaitGroup to wg sync.WaitGroup?**

A deadlock occurs at wg.wait() since wg becomes a copy in each function and the first wg doesn't get marked as done thus leading
to deadlock.

**What happens if you remove the buffer on the channel match?**

At the last name Seek will write to match but without any reader, this is due to the uneven number of persons. 

**What happens if you remove the default-case from the case-statement in the main function?**

The first case will always be true due to the uneven nummber, so removing default won't cause deadlock. The default case is executed if none of the case criterias are matched.
Since the program have 5 people default will be empty, if the program hade even number of people remvoving default would create
a deadlock.

>Task 2 - Fractal Images

**How many CPUs does you program use? How much faster is your parallel version?**

The program uses 4 CPUs. Before adding goroutine the program ran the process on average at 25 sec. After adding
the goroutine it run on average at 11.5 sec

