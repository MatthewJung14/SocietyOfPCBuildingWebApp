describe('RentComponent', () => {
    beforeEach(() => {
      cy.visit('/rent');
    });
  
    it('should reserve a time slot', () => {
      const firstName = 'John';
      const lastName = 'Doe';
      const email = 'johndoe@example.com';
      const timeSlot = '9AM';
  
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
  });
  