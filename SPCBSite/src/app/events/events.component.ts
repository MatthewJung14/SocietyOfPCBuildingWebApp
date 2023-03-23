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
      '../../assets/IMG_0016.JPG',
      imageAlt: 'GBMpic#16',
    },
    {
      imageSrc:
      '../../assets/IMG_0015.JPG',
      imageAlt: 'GBMpic#15',
    },
    {
      imageSrc:
      '../../assets/IMG_0012.JPG',
      imageAlt: 'GBMpic#12',
    },
    {
      imageSrc:
      '../../assets/IMG_0017.JPG',
      imageAlt: 'GBMpic#17',
    },
  ]
}
