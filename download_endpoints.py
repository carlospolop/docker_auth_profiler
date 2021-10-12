import yaml
import requests
import json
import os


def clean_endpoint(endpoint):
    id_example = "71a12335e123"
    img_name_example = "image_name"
    return endpoint.replace("{id}", id_example).replace("{name}", img_name_example)

def main():
    url_endpoints = "https://docs.docker.com/engine/api/v1.40.yaml"
    r = requests.get(url_endpoints)
    y_loaded = yaml.safe_load(r.text)

    endpoints_json = {"endpoints": []}
    for endpoint,val in y_loaded["paths"].items():
        if "get" in val:
            endpoints_json["endpoints"].append({
                "path": clean_endpoint(endpoint),
                "method": "get",
                "summary": val["get"]["summary"]
            })
        
        if "post" in val:
            endpoints_json["endpoints"].append({
                "path": clean_endpoint(endpoint),
                "method": "post",
                "summary": val["post"]["summary"]
            })
    
    file_json = os.path.dirname(os.path.realpath(__file__)) + "/endpoints.json"
    with open(file_json, 'w') as f:
        json.dump(endpoints_json, f)
    
    print(f"endpoints.json ready! There were {len(endpoints_json['endpoints'])} entries loaded.")


if __name__ == "__main__":
    main()