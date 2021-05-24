import { Component, OnInit, Input } from '@angular/core';
import { FilesizePipe } from '../pipes/filesize.pipe';
import {ActivatedRoute} from '@angular/router'
import { TokenServiceService } from '../services/token-service.service';
interface PeriodicElement {
  position: number,
  name: string,
  size: number,
  result: string,
  link:string
}

@Component({
  selector: 'app-table',
  templateUrl: './table.component.html',
  styleUrls: ['./table.component.scss'],
  providers: [ FilesizePipe]
})
export class TableComponent implements OnInit {
  isLogged: boolean;
  email:String="";
  dataSource: PeriodicElement[];
  displayedColumns: string[] = ['position', 'name', 'size', 'result', 'link'];
  constructor(private filesizePipe: FilesizePipe, private route : ActivatedRoute , public tokenService: TokenServiceService) {
    this.dataSource = [
      {position: 1, name: 'Hydrogen', size: 2056, result: 'H', link: "https://safety4sea.com/wp-content/uploads/2019/08/uk-hydrogen-1200x640.png"},
      {position: 2, name: 'Helium', size: 4800, result: 'He', link: "https://physicsworld.com/wp-content/uploads/2019/05/helium.jpg"},
      {position: 3, name: 'Lithium', size: 5631, result: 'Li', link: "https://www.studer-innotec.com/media/image/0/big_16_9/lithium-1-1.png"},
      {position: 4, name: 'Beryllium', size: 1251, result: 'Be', link: "https://cdn.mos.cms.futurecdn.net/Hm3nU5XTogtP98Z6vGJhkW-1200-80.jpg"},
      {position: 5, name: 'Boron', size: 2312, result: 'B', link: "https://owqo93fpiuc4633lp1zthz57-wpengine.netdna-ssl.com/wp-content/uploads/sites/11/2018/12/0817altmedsboron_1268589-860x573.jpg"},
      {position: 6, name: 'Carbon', size: 1252, result: 'C', link: "https://upload.wikimedia.org/wikipedia/commons/f/f0/Graphite-and-diamond-with-scale.jpg"},
      {position: 7, name: 'Nitrogen', size: 3021, result: 'N', link: "https://cdn.mos.cms.futurecdn.net/2ypciwV6a76sb2eisAtFtF-1200-80.jpg"},
      {position: 8, name: 'Oxygen', size: 1024, result: 'O', link: "https://www.thechemicalengineer.com/media/16177/oxygen-symbol.jpg?width=960"},
      {position: 9, name: 'Fluorine', size: 2400, result: 'F', link: "https://upload.wikimedia.org/wikipedia/commons/thumb/3/3c/Fluorite-270246.jpg/235px-Fluorite-270246.jpg"},
      {position: 10, name: 'Neon', size: 2150023, result: 'Ne', link: "https://cdn.britannica.com/90/22390-050-6623F39C/Neon-neon-symbol-square-properties-some-periodic.jpg"}
    ];
    this.isLogged = this.route.snapshot.params.isLogged;
    this.email = this.route.snapshot.params.email;
  }
  ngOnInit(): void{

  }
  logoutEvent(e:Event) {
    this.tokenService.removeToken()
    console.log(localStorage.getItem('token'))
  }

}
