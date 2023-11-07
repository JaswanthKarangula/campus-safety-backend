SELECT D.*
FROM "Users" U
         INNER JOIN "CustomerOfficerMapping" C ON U."user_id" = C."customer_id"
         INNER JOIN "Drones" D ON U."user_id" = D."user_id"
WHERE U."username" = 'customer_username'; -- Replace 'customer_username' with the actual username of the customer you're interested in
