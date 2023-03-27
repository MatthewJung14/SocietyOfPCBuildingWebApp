import { NavbarComponent } from './navbar.component'
import { createOutputSpy } from 'cypress/angular'

describe('NavbarComponent', () => {
  it('mounts', () => {
    cy.mount(NavbarComponent)
  })

  //Check club name text test
  it('stepper should have the title', () => {
    cy.mount(NavbarComponent)
    cy.get('[data-cy=title]').should('have.text', 'The Society of PC Building')
  })

  //Home button test
  it('clicking the home button', () => {
    cy.mount(NavbarComponent, {
      componentProperties: {
        change: createOutputSpy('changeSpy'),
      },
    })
    cy.get('[data-cy=home]').click()
    cy.get('@changeSpy').should('have.been.calledWith', "Home was Clicked")
  })

})