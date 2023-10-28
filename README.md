
# Football Record Analysis


Football Record Analysis is a Data Analytics project as part of the course curriculum, BCCS-9113 Data Analytics. 

The implementation is a pipeline that cleans the input data to make it available for further data analytics processing.

The input data: [Kaggle](https://www.kaggle.com/datasets/davidcariboo/player-scores)




## Data Details

The data contains the following tables

Player(**Player_ID**, Player_Name, Date_of_Birth, Position) \
Club(**Club_ID**, Team_Name, Stadium_Capacity) \
Player_Stats(Player_ID, Current_Club_ID, Season, Yellow_Cards, Red_Cards, Goals, Assists, Minutes_Played, Player_Valuations)




## Pipeline Details

The data first converts the contents of the CSV file to a SQL database data whilst performing aggregations on it. This allows us to view the per-season statistic of a player. 

### Data Cleaning

Mark and Remove Missing Data: We marked missing data in player_valuations for the statistical imputations phase. 

Outlier Analysis: As part of our outlier analysis task, we removed players who played less than 25% of the minutes in said season. 

Statistical Imputation: [Currently in Development] With this, we were able to estimate player valuations using the two-nearest-neighbours approach.

### Data Transformations

We scaled Goals and Assists by a player in season according to the minutes played statistic. This helped us get a more accurate reflection on the attacking output of players.






    
## Data Analysis

We performed some descriptive and statistical analysis on the dataset. 

We primarily answered questions related to Goals and Assists and its relation to the market valuation of a player.

A link to the analysis will be posted soon!
## Authors

- Jay Prasad

