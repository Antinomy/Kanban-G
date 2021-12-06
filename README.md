# Overview
kanban-G is cmd tool to implment kanban in terminal.

It is base markdown file which name by some rule , it is simple and easy to version control by git.

# quick start
download the source by git , enter the folder in terminal , excute cmd :

    ./kanban ./demo/myTasks/

it is looks like this in terminal:

```shell
AutoGit  : 0 / 3
Kan View : UNKNOWN
Today    : 1206
+--------------------------------------+--------------------------------------+--------------------------------------+--------------------------------------+
| Todo [t]                             | Doing [i]                            | Hold [ho]                            | Done [d]                             |
+======================================+======================================+======================================+======================================+
| [t1] AY-H-ProjectA-0531-doSth.md     | [i1] AY-M-ProjectB-0520-doSth.md     | [ho1] AY-H-ProjectC-2020-doSth.md    | [d1] AY-H-ProjectD-0511-doSth.md     |
| [t2] WG-M-ProjectD-0731-doSth.md     | [i2] GLW-H-ProjectB-0621-doSth.md    | [ho2] ZZ-H-ProjectZ-2020-doSth.md    | [d2] GLW-M-ProjectA-0311-doSth.md    |
|                                      | [i3] GLW-L-ProjectC-0621-doSth.md    |                                      | [d3] LL-L-ProjectC-0411-doSth.md     |
+--------------------------------------+--------------------------------------+--------------------------------------+--------------------------------------+

InputCmd $
```



# create task

The task name rule is : 

    OWNER-PRIORITY-PROJECT-DEADLINE-TITTLE.md

    e.g. AY-H-ProjectA-0531-doSth.md


you can create task file by below cmd,

    InputCmd $ create taskname $banPrefix
or

    InputCmd $ c taskname $banPrefix

c is shortcut of create , similar cmd for other cmds.
banPrefix is shortcut of ban, like [t] equal Todo

```shell
# create task ZZ-H-ProjectZ-9999-doSth.md on toda state
InputCmd $ c ZZ-H-ProjectZ-9999-doSth.md
InputCmd $ c ZZ-H-ProjectZ-9999-doSth.md t
```

# exit kanban
you can just type exit / e, to exit the kanban.

    InputCmd $ exit
    InputCmd $ e

# change the task content
change task cmd using changetask / ct

    changetask $taskKey $TaskItem context
    ct $taskKey $TaskItem context

sth like:

```shell
# change task [t1] owner to AY: 
InputCmd $ ct t1 o AY  

# change task [t2] priority to high
InputCmd $ ct t2 i H  

# change task [t3] dealline to June 5 
InputCmd $ ct t3 d 0605  

# change task [t4] project name as [Travel]
InputCmd $ ct t4 j [Travel]  

# change task [t5] tittle  as t5.md
InputCmd $ ct t5 t t5.md
```

# change the task ban
move the task to target ban , using:

    InputCmd $ changeban / cb $taskKey $banPrefix 

e.g.
```shell
# change tast t1 to doing
InputCmd $ cb t1 i  

# change tast t2 to hold
InputCmd $ cb t2 h  

# change tast t3 to done
InputCmd $ cb t3 d  
```
you can alse using shortway to do it.

    InputCmd $ [banKey] $taskKey

banKey means banName or banPrefix

```shell
# change tast t2 to done
InputCmd $ done t1  
InputCmd $ d t2  

# change tast t3 to hold
InputCmd $ hold t3  
InputCmd $ h t4 

```

# view kan on diff view
here using kan to change diff view.

    InputCmd $  k [kan] $TaskItem(i o d j)

e.g.
```shell

# kan by priority ("PRIORITY", "PRI", "PI", "I")
InputCmd $kan priority
InputCmd $k i

# kan by owner ("OWNER", "OWN", "O")
InputCmd $kan owner
InputCmd $k o

# kan by deadline ("DEADLINE", "DL", "D")
InputCmd $kan deadline
InputCmd $k d

# kan by project ("PROJECT", "PRJ", "PJ", "J")
InputCmd $kan project
InputCmd $k j

```

like cmd [k i] , it will looks like:

```shell
AutoGit  : 1 / 3
Kan View : PRIORITY
Today    : 1206
+------+--------------------------------------+--------------------------------------+--------------------------------------+--------------------------------------+
|      | Todo [t]                             | Doing [i]                            | Hold [ho]                            | Done [d]                             |
+======+======================================+======================================+======================================+======================================+
| H    | [t1] AY-H-ProjectA-0531-doSth.md     | [i2] GLW-H-ProjectB-0621-doSth.md    | [ho1] AY-H-ProjectC-2020-doSth.md    | [d1] AY-H-ProjectD-0511-doSth.md     |
|      |                                      |                                      | [ho2] ZZ-H-ProjectZ-2020-doSth.md    |                                      |
+------+--------------------------------------+--------------------------------------+--------------------------------------+--------------------------------------+
| M    | [t2] WG-M-ProjectD-0731-doSth.md     | [i1] AY-M-ProjectB-0520-doSth.md     |                                      | [d2] GLW-M-ProjectA-0311-doSth.md    |
+------+--------------------------------------+--------------------------------------+--------------------------------------+--------------------------------------+
| L    |                                      | [i3] GLW-L-ProjectC-0621-doSth.md    |                                      | [d3] LL-L-ProjectC-0411-doSth.md     |
+------+--------------------------------------+--------------------------------------+--------------------------------------+--------------------------------------+
```


# short mode
short mode is other view for you to keep foucs on your task , it will remove some ban configed on the view
using cmd like

    | InputCmd $ s
    | InputCmd $ short 
    | InputCmd $ shortmode

like hold ban is configed , it will show:

```shell
AutoGit  : 1 / 3
Kan View : PRIORITY
Today    : 1206
+------+--------------------------------------+--------------------------------------+--------------------------------------+
|      | Todo [t]                             | Doing [i]                            | Hold [ho]                            |
+======+======================================+======================================+======================================+
| H    | [t1] AY-H-ProjectA-0531-doSth.md     | [i2] GLW-H-ProjectB-0621-doSth.md    | [ho1] AY-H-ProjectC-2020-doSth.md    |
|      |                                      |                                      | [ho2] ZZ-H-ProjectZ-2020-doSth.md    |
+------+--------------------------------------+--------------------------------------+--------------------------------------+
| M    | [t2] WG-M-ProjectD-0731-doSth.md     | [i1] AY-M-ProjectB-0520-doSth.md     |                                      |
+------+--------------------------------------+--------------------------------------+--------------------------------------+
| L    |                                      | [i3] GLW-L-ProjectC-0621-doSth.md    |                                      |
+------+--------------------------------------+--------------------------------------+--------------------------------------+

InputCmd $
```

# relaod
As the tools is base on file system , if you change the file by other way , tools don't know , you can user reload cmd

    InputCmd $ rekan
    InputCmd $ r



# auto commit & push to git 
all the task and ban are file and folder.
if you won't lose it , best way is version control it by git. 
using cmd:

    InputCmd $ git 
    InputCmd $ g

it had conditions:
1. put myTask folder on your private git repo which sync to your local
2. git acct had configed , able to using git cmd to push 
3. it will auto push when change 3 files


 
# open task
just a way to open md file by your editor 

using cmd:
'''shell
InputCmd $ o [open] $taskKey

 e.g  
InputCmd $ open t1
InputCmd $ o t2
'''

It had conditions:
1. it only tested on osx env.
2. md file had config default editor by osx 

