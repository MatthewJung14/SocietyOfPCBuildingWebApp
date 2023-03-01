describe('template spec', () => {
  it('passes', () => {
    cy.visit('http://localhost:4200/')
  })

  // it('Does not do much!', () => {
  //   expect(true).to.equal(true)
  // })


  it('successfully loads', () => {
    cy.visit('http://localhost:4200/login')
  })
  
})