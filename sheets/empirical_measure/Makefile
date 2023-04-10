# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make empirical_measure.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

empirical_measure.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > empirical_measure.tex

empirical_measure.pdf: ../../*.tex ../../trademark.pdf *.tex empirical_measure.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode empirical_measure.tex
	../../latexrun empirical_measure.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: empirical_measure.pdf
	open empirical_measure.pdf

o:
	make open

terms: empirical_measure.tex
	bbk sheet -mode ts empirical_measure.tex

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
	rm -f empirical_measure.*
	make reset
	bbk sheet -mode mk > Makefile