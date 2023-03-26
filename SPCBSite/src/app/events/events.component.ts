import { Component } from '@angular/core';

@Component({
  selector: 'app-events',
  templateUrl: './events.component.html',
  styleUrls: ['./events.component.css']
})
export class EventsComponent {
  images = [
    {
      imageSrc:
      '../../assets/springgbm1.png',
      imageAlt: 'Spring GBM 1',
    },
    {
      imageSrc:
      '../../assets/springgbm2.png',
      imageAlt: 'Spring GBM 2',
    },
    {
      imageSrc:
      '../../assets/springgbm3.png',
      imageAlt: 'Spring GBM 3',
    },
  ]
}
