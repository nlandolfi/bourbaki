# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make distance_asymmetry.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

distance_asymmetry.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > distance_asymmetry.tex

distance_asymmetry.pdf: ../../*.tex ../../trademark.pdf *.tex distance_asymmetry.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode distance_asymmetry.tex
	../../latexrun distance_asymmetry.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: distance_asymmetry.pdf
	open distance_asymmetry.pdf

o:
	make open

terms: distance_asymmetry.tex
	bbk sheet -mode ts distance_asymmetry.tex

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
	rm -f distance_asymmetry.*
	make reset
	bbk sheet -mode mk > Makefile