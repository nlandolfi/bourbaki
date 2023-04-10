# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make almost_sure_events.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

almost_sure_events.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > almost_sure_events.tex

almost_sure_events.pdf: ../../*.tex ../../trademark.pdf *.tex almost_sure_events.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode almost_sure_events.tex
	../../latexrun almost_sure_events.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: almost_sure_events.pdf
	open almost_sure_events.pdf

o:
	make open

terms: almost_sure_events.tex
	bbk sheet -mode ts almost_sure_events.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f almost_sure_events.*
	make reset
	bbk sheet -mode mk > Makefile