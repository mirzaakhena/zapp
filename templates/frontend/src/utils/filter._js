import numeral from 'numeral'
import moment from 'moment'

export function digitGrouping(x) {
  if (x === '') {
    return ''
  }
  return numeral(x).format('0,0')  
}

export function dateFormat(value) {
  if (value) {
    return moment(String(value)).format('YYYY-MM-DD hh:mm')
  }
}
