# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make data_models.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

data_models.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > data_models.tex

data_models.pdf: ../../*.tex ../../trademark.pdf *.tex data_models.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode data_models.tex
	../../latexrun data_models.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: data_models.pdf
	open data_models.pdf

o:
	make open

terms: data_models.tex
	bbk sheet -mode ts data_models.tex

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
	rm -f data_models.*
	make reset
	bbk sheet -mode mk > Makefile