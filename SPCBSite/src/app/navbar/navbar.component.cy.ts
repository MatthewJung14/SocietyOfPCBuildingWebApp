import { NavbarComponent } from './navbar.component'

describe('NavbarComponent', () => {
  it('mounts', () => {
    cy.mount(NavbarComponent)
  })
  
  it('stepper should have the home button', () => {
    cy.mount(NavbarComponent)
    cy.get('[data-cy=counter]').should('have.text', 'The Society of PC Building')
  })
})