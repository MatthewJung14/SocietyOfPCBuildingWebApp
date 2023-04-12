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
    cy.get('[name^=email]').type('Matthewjung14@gmail.com')
    cy.get('[name^=password]').type('hello')
    cy.get('[name^=signup]').click()
  })

  it('login test', () => {
    cy.visit('http://localhost:3200')
    cy.contains('LOGIN').click()
    cy.get('[name^=email]').type('matthewjung@ufl.edu')
    cy.get('[name^=password]').type('hel1o')
    cy.get('[name^=login]').click()
    cy.contains('What is')
    cy.contains('SETTINGS')
  })

  it('settings test', () => {
    cy.visit('http://localhost:3200')
    cy.contains('LOGIN').click()
    cy.get('[name^=email]').type('matthewjung@ufl.edu')
    cy.get('[name^=password]').type('hel1o')
    cy.get('[name^=login]').click()
    cy.contains('SETTINGS').click()
    cy.contains('LOG OUT')
    cy.contains('NAME').click()
    cy.contains('SAVE CHANGES')
    cy.contains('GO BACK').click()
    cy.contains('PASSWORD').click()
    cy.contains('GO BACK').click()
    cy.contains('LOG OUT').click()
  })

  it('events test', () => {
    cy.visit('http://localhost:3200')
    cy.contains('EVENTS').click()
    cy.get('[name^=carousel]')
  })
  
})