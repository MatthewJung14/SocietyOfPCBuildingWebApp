import { Component, Input, OnInit} from '@angular/core';

interface SPCBImages {
  imageSrc: string;
  imageAlt: string;
}

@Component({
  selector: 'app-caraousel',
  templateUrl: './caraousel.component.html',
  styleUrls: ['./caraousel.component.scss']
})
export class CaraouselComponent implements OnInit{

  @Input() images: SPCBImages[] = []
  @Input() indicators = true;
  @Input() controls = true;

  selectedIndex = 0;

  constructor() { }

  ngOnInit(): void { 
  }

  selectImage(index: number): void {
    this.selectedIndex=index;
  }
}
