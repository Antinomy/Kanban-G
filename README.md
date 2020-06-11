# Kanban-G
Kanban Project reborn by Go


# exit
Input Cmd $ exit / e

# relaod
Input Cmd $ rekan / r

# kan
Input Cmd $  kan / k [TaskItem]  
[TaskItem] = i / o / j /d   
i = case "PRIORITY", "PRI", "PI", "I"  
o = case "OWNER", "OWN", "O"  
d = case "DEADLINE", "DL", "D"  
j = case "PROJECT", "PRJ", "PJ", "J"  

# kan short mode
Input Cmd $ shortmode / short / s

# change task
Input Cmd $  changetask / ct [taskKey] [TaskItem] context  
e.g.   
ct t1 o AY  
ct t2 i H  
ct t3 d 0605  
ct t4 j [Travel]  

# change ban
## way1:
Input Cmd $ changeban / cb [taskKey] [banPrefix]  
e.g.  
cb t1 i  
cb t2 h  
cb t3 d  

## way2:
Input Cmd $ [banKey] [taskKey]
banKey = banName or banPrefix  
e.g  
done t1  
d t2  
hold t3  
h t4 


# create ban task
Input Cmd $ create / c taskname [banPrefix]  
e.g.  
c ZZ-H-ProjectZ-9999-doSth.md   
c ZZ-H-ProjectZ-9999-doSth.md t  
 
# open task 
 Input Cmd $ open / o [taskKey]
 e.g  
 open t1
 o t2