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

  it('test settings page', () => {
    cy.visit('http://localhost:3200')
    cy.get('[data-cy=settings]').click()
    cy.contains('LOG OUT')
    cy.contains('NAME').click()
    cy.contains('SAVE CHANGES')
    cy.contains('GO BACK').click()
    cy.contains('PASSWORD').click()
    cy.contains('GO BACK').click()
    cy.contains('LOG OUT').click()
    
  })
  
})