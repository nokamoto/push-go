package nokamoto.github.com.pushgoandroid;

import android.util.Log;

import io.grpc.ManagedChannel;
import io.grpc.okhttp.OkHttpChannelBuilder;

class GrpcUtils {
    private final static String TAG = GrpcUtils.class.getSimpleName();

    private final static String HOST = BuildConfig.GRPC_HOST;
    private final static int PORT = Integer.parseInt(BuildConfig.GRPC_PORT);
    private final static String APIKEY = BuildConfig.GRPC_APIKEY;

    private final static String LOG_HOST = BuildConfig.GRPC_LOG_HOST;
    private final static int LOG_PORT = Integer.parseInt(BuildConfig.GRPC_LOG_PORT);
    private final static String LOG_APIKEY = BuildConfig.GRPC_LOG_APIKEY;

    public static ManagedChannel newChannel(String host, int port, String apikey) {
        Log.i(TAG, String.format("grpc channel: host=%s, port=%d", host, port));

        return OkHttpChannelBuilder.
                forAddress(host, port).
                intercept(new GrpcApiKeyInterceptor(apikey)).
                usePlaintext(true).
                build();
    }

    public static ManagedChannel newChannel() {
        return newChannel(HOST, PORT, APIKEY);
    }

    public static ManagedChannel newLogChannel() {
        return newChannel(LOG_HOST, LOG_PORT, LOG_APIKEY);
    }
}