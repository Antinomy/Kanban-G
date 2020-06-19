
# exit                 | InputCmd $ e [exit]
# relaod               | InputCmd $ r [rekan]
# short mode           | InputCmd $ s [short / shortmode]
# commit & push to git | InputCmd $ g [git] 
# kan                  | InputCmd $  k [kan] $TaskItem(i o d j)  
# k i                  | case "PRIORITY", "PRI", "PI", "I"  
# k o                  | case "OWNER", "OWN", "O"  
# k d                  | case "DEADLINE", "DL", "D"  
# k j                  | case "PROJECT", "PRJ", "PJ", "J"  

# change task          | InputCmd $  changetask / ct $taskKey $TaskItem context  
e.g.   
ct t1 o AY  
ct t2 i H  
ct t3 d 0605  
ct t4 j [Travel]  

# change ban  ## way1: | InputCmd $ changeban / cb $taskKey $banPrefix 
e.g.  
cb t1 i  
cb t2 h  
cb t3 d  

#              ##way2: | InputCmd $ [banKey] $taskKey
banKey = banName or banPrefix  
e.g  
done t1  
d t2  
hold t3  
h t4 

# create ban task      | InputCmd $ c [create] taskname $banPrefix   
e.g.  
c ZZ-H-ProjectZ-9999-doSth.md   
c ZZ-H-ProjectZ-9999-doSth.md t  
 
# open task            | InputCmd $ o [open] $taskKey
 e.g  
 open t1
 o t2