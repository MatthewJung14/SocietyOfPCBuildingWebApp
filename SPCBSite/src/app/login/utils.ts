export function stringifyWithZone(obj: any): string {
    const zoneState = obj?.__zone_symbol__state;
    const zoneValue = obj?.__zone_symbol__value;
    const json = JSON.stringify(obj);
  
    if (zoneState !== null && typeof zoneValue !== 'undefined') {
      return json.replace('}', `,"__zone_symbol__state":${zoneState},"__zone_symbol__value":${JSON.stringify(zoneValue)}}`);
    }
  
    return json;
  }

  export interface ZoneObject {
    __zone_symbol__state?: boolean;
    [key: string]: any;
    __zone_symbol__value: any;
  }