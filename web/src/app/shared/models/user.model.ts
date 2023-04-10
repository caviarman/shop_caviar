import { Order } from './order.model'; 

export class User {
    id!: number;
    telegram_id!: number;
    name!: string;
    telegram_username!: string;
    queryId!: string;
    authDate!: number;
    hash!: string;
    orders!: Order[];
  
    constructor(data: string) {
      this.parseData(data);
    }
  
    private parseData(data: string): void {
      const pairs = data.split('&');
      for (let i = 0; i < pairs.length; i++) {
        const values = pairs[i].split('=');
        switch (values[0]) {
          case 'user':
            const userData = JSON.parse(decodeURIComponent(values[1]));
            this.telegram_id = userData.id;
            this.name = userData.first_name + ' ' + userData.last_name;
            this.telegram_username = userData.username;
            
            break;
          case 'query_id':
            this.queryId = values[1];
            break;
          case 'auth_date':
            this.authDate = parseInt(values[1], 10);
            break;
          case 'hash':
            this.hash = values[1];
            break;
          default:
            break;
        }
      }
    }
  }
  