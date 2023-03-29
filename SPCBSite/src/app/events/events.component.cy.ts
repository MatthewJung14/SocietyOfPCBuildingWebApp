import { EventsComponent } from './events.component';

describe('EventsComponent', () => {
  it('mounts', () => {
    cy.mount(EventsComponent)
    cy.get('[name^=carousel]')
  })
})