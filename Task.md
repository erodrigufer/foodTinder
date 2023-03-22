## Summary
Your task is to create a food "Tinder". 
You can use any technology you want for your REST backend.
 
## Context
Users of a web app should be able to vote on products that the web app receives from an endpoint from our systems in a Tinder-like interface and user experience. 
The web app receives the products to vote from from this endpoint:
https://amperoid.tenants.foodji.io/machines/4bf115ee-303a-4089-a3ea-f6e7aae0ab94
  
## Details:
In this challenge, you only write the backend for handling product votes. 
It should be able to:
	1. Generate a unique session id
	2. Store/Update product votes for a given session id
	3. Retrieve existing votes for products for a given session id
	4. Retrieve aggregated average scores for products across all session ids
The stored votes and sessions should be persisted in a database of your choice.
    
For extra credit deploy your backend on a (temporary) cloud infrastructure of your choice.
	 
What we are looking for:
	- Well setup project using best practices of the selected framework/library
	- Well written and documented code
	- Error resilient design
