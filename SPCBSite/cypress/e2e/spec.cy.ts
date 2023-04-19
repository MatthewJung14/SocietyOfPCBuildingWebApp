describe('template spec', () => {
  //Visit localhost test
  it('passes', () => {
    cy.visit('http://localhost:3200/')
  })

  // it('Does not do much!', () => {
  //   expect(true).to.equal(true)
  // })

  //Find text in the SPCB web app test
  it('successfully loads', () => {
    cy.visit('http://localhost:3200')
    cy.contains('HOME').click()
    cy.contains('OFFICERS').click()
    cy.contains('RENT A PC').click()
    cy.contains('UPCOMING EVENTS').click()
  })

  it('signup test', () => {
    cy.visit('http://localhost:3200')
    cy.contains('SIGN UP').click()
    cy.get('[name^=fname]').type('Matthew')
    cy.get('[name^=lname]').type('Jung')
    cy.get('[name^=email]').type('matthewjung@ufl.edu')
    cy.get('[name^=password]').type('hello')
    cy.get('[name^=signup]').click()
    cy.contains('Forgot Password?')
  })

  it('login test', () => {
    cy.visit('http://localhost:3200')
    cy.contains('LOGIN').click()
    cy.get('[name^=email]').type('matthewjung@ufl.edu')
    cy.get('[name^=password]').type('hello')
  })

  it('settings test', () => {
    cy.visit('http://localhost:3200/settings')
    cy.contains('LOG OUT')
    cy.contains('NAME').click()
    cy.contains('SAVE CHANGES')
    cy.contains('GO BACK').click()
    cy.contains('PASSWORD').click()
    cy.contains('GO BACK').click()
    cy.contains('LOG OUT').click()
    cy.contains('LOG IN')
  })

  it('events test', () => {
    cy.visit('http://localhost:3200')
    cy.contains('EVENTS').click()
    cy.get('[name^=carousel]')
  })

  it('signup navigates to login', () => {
    cy.visit('http://localhost:3200')
    cy.contains('SIGN UP').click()
    cy.get('[name^=fname]').type('Matthew')
    cy.get('[name^=lname]').type('Jung')
    cy.get('[name^=email]').type('test14@gmail.com')
    cy.get('[name^=password]').type('test')
    cy.get('[name^=signup]').click()
  })

  it('settings navigates to home', () => {
    cy.visit('http://localhost:3200')
    cy.contains('LOGIN').click()
    cy.get('[name^=email]').type('test14@gmail.com')
    cy.get('[name^=password]').type('test')
    cy.get('[name^=login]').click()
    cy.contains('SETTINGS').click()
    cy.contains('LOG OUT')
    cy.get('[name^=logOut]').click()
    cy.url().should('eq', 'http://localhost:3200/home')
  })

  it('should reserve a time slot', () => {
    const firstName = 'John';
    const lastName = 'Doe';
    const email = 'johndoe@example.com';
    const timeSlot = '9AM';

    cy.visit('http://localhost:3200');
    cy.contains('RENT A PC').click()
    cy.get('input[name="firstName"]').type(firstName);
    cy.get('input[name="lastName"]').type(lastName);
    cy.get('input[name="email"]').type(email);
    cy.get('input[name="timeSlot"]').type(timeSlot);

    cy.get('button').contains('SUBMIT TIME').click();

    cy.get('.string-array__item').contains(timeSlot).should('not.exist');
    cy.get('.string-array__item').contains(`Reserved by: ${firstName} ${lastName} (${email})`);
  });

  it('should prevent reserving a time slot with the same email', () => {
    const firstName1 = 'John';
    const lastName1 = 'Doe';
    const email1 = 'johndoe@example.com';
    const timeSlot1 = '9AM';
    const firstName2 = 'Jane';
    const lastName2 = 'Doe';
    const email2 = 'janedoe@example.com';
    const timeSlot2 = '10AM';

    cy.visit('http://localhost:3200');
    cy.contains('RENT A PC').click()
    cy.get('input[name="firstName"]').type(firstName1);
    cy.get('input[name="lastName"]').type(lastName1);
    cy.get('input[name="email"]').type(email1);
    cy.get('input[name="timeSlot"]').type(timeSlot1);
    cy.get('button').contains('SUBMIT TIME').click();

    cy.get('input[name="firstName"]').type(firstName2);
    cy.get('input[name="lastName"]').type(lastName2);
    cy.get('input[name="email"]').type(email1);
    cy.get('input[name="timeSlot"]').type(timeSlot2);
    cy.get('button').contains('SUBMIT TIME').click();

    cy.get('.string-array__item').contains(timeSlot1).should('not.exist');
    cy.get('.string-array__item').contains(`Reserved by: ${firstName1} ${lastName1} (${email1})`);
    cy.get('.string-array__item').contains(timeSlot2).should('exist').and('not.contain', email1);
  });
  
})