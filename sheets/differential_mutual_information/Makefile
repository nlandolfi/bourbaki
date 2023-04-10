# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make differential_mutual_information.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

differential_mutual_information.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > differential_mutual_information.tex

differential_mutual_information.pdf: ../../*.tex ../../trademark.pdf *.tex differential_mutual_information.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode differential_mutual_information.tex
	../../latexrun differential_mutual_information.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: differential_mutual_information.pdf
	open differential_mutual_information.pdf

o:
	make open

terms: differential_mutual_information.tex
	bbk sheet -mode ts differential_mutual_information.tex

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
	rm -f differential_mutual_information.*
	make reset
	bbk sheet -mode mk > Makefile