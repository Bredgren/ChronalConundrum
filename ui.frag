precision mediump float;

varying highp vec2 vUV;

uniform sampler2D uTexture;
uniform vec4 uRect;

void main(void) {
	// float x = vUV.s / uRect[2] + uRect.x;
	// float y = vUV.t / uRect[3] + uRect.y;
	vec2 pos = vUV.st * uRect.zw + uRect.xy;
	gl_FragColor = texture2D(uTexture, pos);
}
