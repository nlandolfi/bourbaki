# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make feature_maps.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

feature_maps.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > feature_maps.tex

feature_maps.pdf: ../../*.tex ../../trademark.pdf *.tex feature_maps.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode feature_maps.tex
	../../latexrun feature_maps.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: feature_maps.pdf
	open feature_maps.pdf

o:
	make open

terms: feature_maps.tex
	bbk sheet -mode ts feature_maps.tex

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
	rm -f feature_maps.*
	make reset
	bbk sheet -mode mk > Makefile