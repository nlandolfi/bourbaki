# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make confusion_matrices.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

confusion_matrices.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > confusion_matrices.tex

confusion_matrices.pdf: ../../*.tex ../../trademark.pdf *.tex confusion_matrices.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode confusion_matrices.tex
	../../latexrun confusion_matrices.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: confusion_matrices.pdf
	open confusion_matrices.pdf

o:
	make open

terms: confusion_matrices.tex
	bbk sheet -mode ts confusion_matrices.tex

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
	rm -f confusion_matrices.*
	make reset
	bbk sheet -mode mk > Makefile