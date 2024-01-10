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

## Architecture
<br> <br>
    ![image](https://github.com/Samridhi16/fetch-rewards/assets/26019260/89056cfe-74bc-4d57-ac1f-7329d609199b)
    
## Demo (Click on the thumbnails below to play the video)
1. **Endpoint: Process Receipts**
    - **Path:** `/receipts/process`
    - **Method:** POST
    - **Payload:** Receipt JSON
    - **Response:** JSON containing an ID for the receipt.
<br> <br>
    [![POST](https://github.com/Samridhi16/fetch-rewards/assets/26019260/133fff30-7e96-4fa3-8cf6-7efa005ef65e)](https://github.com/Samridhi16/fetch-rewards/assets/26019260/eccfb44c-b1d1-47ca-ada9-69b942b23e2a)

2. **Endpoint: Get Points**
    - **Path:** `/receipts/{id}/points`
    - **Method:** GET
    - **Response:** A JSON object containing the number of points awarded.
<br> <br>
    [![GET](https://github.com/Samridhi16/fetch-rewards/assets/26019260/8f6f0915-3fcf-47e0-967d-6a7aabd62263)](https://github.com/Samridhi16/fetch-rewards/assets/26019260/a4994976-155d-453b-ab85-f5f5aa41df3c)

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
To run the project:
- Open the project fetch-rewards in Visual Studio Code or any editor of your choice.
- Open the terminal and import all the necessary packages like gin to your project directory.
- Run the following command and test the GET and POST methods using Postman or a web browser on localhost:8080
    ```sh
    go run main.go
    ```

    




  


