# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make classifier_errors.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

classifier_errors.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > classifier_errors.tex

classifier_errors.pdf: ../../*.tex ../../trademark.pdf *.tex classifier_errors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode classifier_errors.tex
	../../latexrun classifier_errors.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: classifier_errors.pdf
	open classifier_errors.pdf

o:
	make open

terms: classifier_errors.tex
	bbk sheet -mode ts classifier_errors.tex

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
	rm -f classifier_errors.*
	make reset
	bbk sheet -mode mk > Makefile