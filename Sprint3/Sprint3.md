**Work completed for the frontend**
- Created a settings page where users can choose to edit their first and last name, their password, or their email. On either page the user would type in their updated name, password, or email, and hit "SAVE CHANGES". If they want to go back to the previous menu, they can just click on the "GO BACK" button. 
- If the user wants to log out, they can click on the settings page, then click "LOG OUT" on the bottom right corner of the of the page.
- Changed the favicon to have our club's logo on it.
- The Navbar now changes based on whether the user is logged in or not. 
- The login button navigates to the home page if the user successfully logs in.
- The signup button navigates to the login page if the user successfully signs up.
- Created a new page called "UPCOMING EVENTS". The Society of PC Building's upcoming events will be posted here. 
- Added more images of various officers on the officer's page.

**List frontend tests**
- Test settings page: Wrote an E2E test that tests the functionality of the settings button. This includes being able to navigate properly to the change name and change password sections, being able to change the user's name in the database, being able to change the user password in the database, and being able to log off.
- 
**Work completed for the backend**
- Created a update user firstname and lastname function.
- Created a update user email function.
- Implemented functionality to be able to send emails to requested users located within the database.
- adjusted some functions to be able to accessiable to the front end.

**List Backend tests**
- we tested the sendemail function to be able to send emails to a specific email located within the database.
- created a update user test case for first and last name.
- created a update user email function.
