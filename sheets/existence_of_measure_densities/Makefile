# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make existence_of_measure_densities.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

existence_of_measure_densities.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > existence_of_measure_densities.tex

existence_of_measure_densities.pdf: ../../*.tex ../../trademark.pdf *.tex existence_of_measure_densities.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode existence_of_measure_densities.tex
	../../latexrun existence_of_measure_densities.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: existence_of_measure_densities.pdf
	open existence_of_measure_densities.pdf

o:
	make open

terms: existence_of_measure_densities.tex
	bbk sheet -mode ts existence_of_measure_densities.tex

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
	rm -f existence_of_measure_densities.*
	make reset
	bbk sheet -mode mk > Makefile