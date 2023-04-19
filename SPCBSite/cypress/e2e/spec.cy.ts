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
    cy.contains('LOGIN')
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
    cy.visit('http://localhost:3200/settings')
    cy.contains('LOG OUT')
    cy.get('[name^=logOut]').click()
    cy.url().should('eq', 'http://localhost:3200/home')
  })

  describe('RentComponent', () => {
    beforeEach(() => {
      cy.visit('http://localhost:3200');
      cy.contains('RENT A PC').click()
    });
  
    it('should have form fields', () => {
      cy.get('input[name="firstName"]').should('exist');
      cy.get('input[name="lastName"]').should('exist');
      cy.get('input[name="email"]').should('exist');
      cy.get('input[name="timeSlot"]').type('11PM', { force: true });
    });
  
    it('should submit form with valid data', () => {
      cy.get('input[name="firstName"]').type('John');
      cy.get('input[name="lastName"]').type('Doe');
      cy.get('input[name="email"]').type('johndoe@example.com');
      cy.get('input[name="timeSlot"]').type('8AM', { force: true });
      cy.get('button').contains('SUBMIT TIME').click();
    });
  
    it('should not submit form with invalid data', () => {
      cy.get('input[name="firstName"]').type('John');
      cy.get('input[name="lastName"]').type('Doe');
      cy.get('input[name="email"]').type('johndoe@example.com');
      cy.get('button').contains('SUBMIT TIME').click();
      cy.get('.string-array__item').contains('johndoe@example.com').should('not.exist');
      cy.url().should('not.include', '/home');
    });
  });
  
})