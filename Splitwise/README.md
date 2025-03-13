# Building Splitwise

## Goals

- Design app that helps tracks expenses
    - Either in a group or a set of individuals

## Requirements

- Add expense
- Edit expense
- Settle expense
- Add group
    - Add/Edit expense
    - Settle expense (Simplified)
- Optional
    - Comments
    - Activity Log

## Defining objects and behaviors

- Behavior driven approach
    - What the user needs, and how my system needs to behave to satisfy our requirements
- State based approach 
    - ER Diagram approach
        - Design the DB schema, and then work on the API design, based on this

## ER Diagram approach

- Different people paid different amounts
- Expense: 
    - Id, GroupId
    - Map<user,balance>
    - Metadata:
        - Timestamp
        - Image
        - Comments
    - Settled (boolean)

- User
    - Id, ImageUrl, Bio

- Balance
    - Integers
    - Currency

- Group
    - Id, ImageUrl, Description, Title
    - Users List

## Balances algo

- We are building a payment graph
    - Nodes represent the users
    - Connection represents call for money
    - We want to minise the number of edges and transactions
    - Initially we need to take pools of positive and negative numbers
        - Elimininate one by one, by taking the largest from each side
        - Heap allows us to do this
        - Everytime topmost node is going to be pulled out of the heap
            - Whichever node survives, is pushed back to the heap
            - Quick optimisation would be to cancel equal numbers in the 2 heaps
        - Alternatively, SubsetSum problem can be used here

## APIs and coding requirements

- Tackling flow:
    - APIs -> Caching -> Concurrency -> Testing

- APIs
    - `getUser`, `getGroup`, `getExpense`
    - These are simple CRUD APIs
    - We can skip coding simple bits
    - We are expected to code algorithmicly complex bits
        - Or things with nances like Concurreny or Caching
    
- Caching
    - We will cache things that are accessed by many people
    - `getGroupExpenses`, `getGroupPaymentGraph`

- Concurrency
    - Make object immutable
    - Immutable objects are giving us eventual consistency
    - `getGroupExpense` and `getGroupPaymentGraph` needs to have immutable objects
        - These can be cached
    - We return old stale value, not inconsistent value

- Testing
    - Unit testing
    - We will check if our payment graph is created correctly

## Architecture

- Final step before we write code (Similar to HLD)

## Code

- We are expected to define objects, and their behavior
    - Easier to solve
    - Make sure they are named well
    - This is the most important bit

- Composition
    - Break down large objects into smaller object
        - Each small object could be complex in itself
    - Overall complexity is reduced when we abstract them out

- Inheritance
    - Be careful of inheritances
    - They should be directly from interfaces

- Algorithm
    - How do we create a payment graph

- Test cases
    -  

## Service

- GroupService
    - Need payment graph
    - Or all expenses related to the group
    - `GetGroupPaymentGraph` is summation of all group expenses + paymentGraph
        - First we need to sum up all the balances (per user)
        - `select * from expenses where group_id='group_id'`
    - Nested array or map => requires simplification (Abstraction)
- Mock the bit where we need DB connection 

- ExpenseService