# fetch-rewards

Design Decisions:
1) Unique ID associated with receipt created using receipt details in order to make sure duplicate entries are not recorded and same entries are not recorded multiple times.
2) Trade off for whether to compute rewards points when receipts are posted VS when points are fetched
    a) Compute points when receipts are posted:
        If stored in memory at the time of POST, any changes in points calculation logic will affect all the records and we will need to modify all the records.
    b) Compute points when points are fetched:
        If the points computation is very heavy(rare case), it might take a while to fetch the points, but when the points calculation logic is changed, it is easy to make a GET request to modify the records.
    The implemented project uses second technique (b) as it helps us to utilize our resources efficiently.
     