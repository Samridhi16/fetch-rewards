# Fetch Rewards

## Design Decisions

1. **Unique ID associated with receipt**: 
    - Created using receipt details to ensure that the same entries are not recorded multiple times.

2. **Trade-off for when to compute rewards points**: Whether to compute points when receipts are posted or when points are fetched.
    - **Compute points when receipts are posted**: 
        - If stored in memory at the time of POST, any changes in points calculation logic will affect all the records. As POST operations for each entry are done only once, we will need to modify all the previously stored records.
    - **Compute points when points are fetched**: 
        - If the points computation is very heavy, it might take a while to fetch the points. However, this is a rare case.
        - When the points calculation logic is changed, it is easier to make GET requests to modify the records.
    - The implemented project uses the second technique (Compute points when points are fetched) as it helps us to utilize our resources efficiently.

## Demo
1. **Endpoint: Process Receipts**
    - **Path:** `/receipts/process`
    - **Method:** POST
    - **Payload:** Receipt JSON
    - **Response:** JSON containing an ID for the receipt.

2. **Endpoint: Get Points**
    - **Path:** `/receipts/{id}/points`
    - **Method:** GET
    - **Response:** A JSON object containing the number of points awarded.

## How to run:

### Prerequisites
- Ensure you have [Go](https://golang.org/dl/) installed.
- Any IDE like [Visual Studio Code](https://code.visualstudio.com/download) is recommended.

### Clone the Repository
```sh
git clone https://github.com/Samridhi16/fetch-rewards.git
cd fetch-rewards
```
### Running the Project
To run the project, execute the following command:
    ```sh
    go run main. go
     ```




  


