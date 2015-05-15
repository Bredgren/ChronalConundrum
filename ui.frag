precision mediump float;

varying highp vec2 vTextureCoord;

uniform sampler2D uSampler;
uniform vec4 uRect;

void main(void) {
	// float x = vTextureCoord.s / uRect[2] + uRect.x;
	// float y = vTextureCoord.t / uRect[3] + uRect.y;
	vec2 pos = vTextureCoord.st * uRect.zw + uRect.xy;
	gl_FragColor = texture2D(uSampler, pos);
}
