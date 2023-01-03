## Trainer Source of Truth
### MVP TODO
#### Data + CRUD Endpoints
- [] Trainers
- [] Customers
- [] Accounts Receivable
- [] Payments
- [] Customer Files
#### UI 
- [] Login
- [] Customer
- [] Accounts Receivable
- [] Payments
- [] Customer Files

#### Unique functionality
- When creating payments, (as a trainer), a list of accounts receivable should be selected that the payment will be applied to. If the payment is less than the amount of the first receivable, then it will continue to have a balance due. If the payment is more, it will go on to apply to the next receivable in the list. If there is a remaining amount from the payment after paying off all the receivables in the list, (insert some desireable functionality here...).