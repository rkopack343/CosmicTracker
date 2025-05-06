# CosmicTracker
Cosmic Tracker for FFXIV

Idea: Scrape the cosmic exploration report website - https://na.finalfantasyxiv.com/lodestone/cosmic_exploration/report/ - and store data about the cosmic exploration over time.
This data is updated every 30 minutes so we can ping the website every ~25 minutes and it'll be okay 

## Notable Selectors

- #contents > div > div.cosmic__contents > div > div.cosmic__report__header > p.cosmic__report__update
  - This has the Last Updated in a span with a random id like 'datetime-[hash]'

- #Aether, #Crystal, #Dynamis, #Primal
  - The main divs which hold all cosmic report information (class="cosmic__report__dc")

- Under each DC div is a span and a div, the div is class "cosmic__report__world" and "show"
   - Under that div is a separate div for each server on that DC.
   - Progress is under something like #Aether > div > div:nth-child(1) > div.cosmic__report__status > div  (span, class has a 'gauge-#')
      - If no gauge it's full
   - The Stage is under something like #Aether > div > div:nth-child(1) > div.cosmic__report__grade > div


## Storage

Table with following columns:

- UUID
- DataCenter (Not really needed but sure)
- Server Name
- Gauge Progress
- Grade Number
- Timestamp reported on Page
- Our own Timestamp


## Logic

### Retrieval & Process

1. Scrape the website and put all of the data into an array of objects representing our storage minus UUID/Timestamps
1. For each object...
   1. Check to see if Grade Number > Stored Grade Number. If so, insert new Object.
   1. Check to see if Gauge Progress > Stored Gauge Progress. If so, insert new Object.

### Display

Need some kind of chart program that shows progress...  
Would be nice to mark grade bumps with a star or something and display the datetime on it.  
Gauge Progress always seems to be 8 steps?


### FAQ

Q. Why Go?  
A. I wanted to learn it.