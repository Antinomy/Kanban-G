# Kanban-G
Kanban Project reborn by Go


# exit
> exit / e

# relaod
> rekan / r

# kan
>  kan / k [TaskItem]  
[TaskItem] = i / o / j /d   
i = case "PRIORITY", "PRI", "PI", "I"  
o = case "OWNER", "OWN", "O"  
d = case "DEADLINE", "DL", "D"  
j = case "PROJECT", "PRJ", "PJ", "J"  

# kan short mode
> k [TaskItem] short / s

# change task
>  changetask / ct [targetKey] [TaskItem] context  
e.g.   
ct t1 o AY  
ct t2 i H  
ct t3 d 0605  
ct t4 j [Travel]  

# change ban
> changeban / cb [targetKey] [banPrefix]  
e.g.  
cb t1 i  
cb t2 h  
cb t3 d  


 