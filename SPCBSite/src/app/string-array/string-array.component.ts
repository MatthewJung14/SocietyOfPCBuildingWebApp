import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-string-array',
  template: `
    <mat-card>
      <mat-card-content>
        <mat-list>
          <mat-list-item *ngFor="let str of strings">{{ str }}</mat-list-item>
        </mat-list>
      </mat-card-content>
    </mat-card>
  `,
  styleUrls: ['./string-array.component.css'] // Add this line

})
export class StringArrayComponent {
  @Input() strings: string[];

  constructor() {
    this.strings = []; // Initializes the strings property in the constructor
  }
}
