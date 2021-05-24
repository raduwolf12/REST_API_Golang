import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'filesize'
})
export class FilesizePipe implements PipeTransform {
  private units = ['bytes', 'KB', 'MB', 'GB', 'TB', 'PB'];

  transform(size: number) {
    if (isNaN(parseFloat(String(size))) || !isFinite(size)) return '?';

        let unit = 0;

        while (size >= 1024) {
            size /= 1024;
            unit++;
        }
    return size.toFixed(2) + ' ' + this.units[unit];
  }

}
