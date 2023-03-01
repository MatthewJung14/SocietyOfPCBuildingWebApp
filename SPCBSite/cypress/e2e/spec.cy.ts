describe('template spec', () => {
  it('passes', () => {
    cy.visit('http://localhost:4200/')
  })

  // it('Does not do much!', () => {
  //   expect(true).to.equal(true)
  // })


  it('successfully loads', () => {
    cy.visit('http://localhost:4200/login')
    cy.contains('HOME').click()
    cy.contains('OFFICERS').click()
    cy.contains('RENT A PC').click()
    cy.contains('LOGIN').click()
    cy.contains('SIGN UP').click()
    cy.contains('First Name')
    cy.contains('Last Name')
    cy.contains('Email')
    cy.contains('Password')
  })
  
})