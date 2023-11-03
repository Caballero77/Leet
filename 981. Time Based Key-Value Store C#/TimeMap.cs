public class TimeMap {

    private struct TimeValue {
        public int Timestamp { get; set; }
        public string Value { get; set; }
    }

    private readonly Dictionary<string, List<TimeValue>> store = new();
    
    public void Set(string key, string value, int timestamp) {
        if (store.TryGetValue(key, out var values)) {
            values.Add(new(){ Timestamp = timestamp, Value = value});
        } else {
            store.Add(key, new() { new(){ Timestamp = timestamp, Value = value}});
        }
    }
    
    public string Get(string key, int timestamp) {
        if (!store.TryGetValue(key, out var values)) {
            return "";
        }

        var start = 0;
        var end = values.Count()-1;
        string result = string.Empty;
        while (start <= end) {
            var middle = (start + end)/2;

            if (values[middle].Timestamp > timestamp) {
                end = middle - 1;
            } else {
                result = values[middle].Value;
                start = middle + 1;
            }
        }

        return result;
    }
}