# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make signed_set_decomposition_existence.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

signed_set_decomposition_existence.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > signed_set_decomposition_existence.tex

signed_set_decomposition_existence.pdf: ../../*.tex ../../trademark.pdf *.tex signed_set_decomposition_existence.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode signed_set_decomposition_existence.tex
	../../latexrun signed_set_decomposition_existence.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: signed_set_decomposition_existence.pdf
	open signed_set_decomposition_existence.pdf

o:
	make open

terms: signed_set_decomposition_existence.tex
	bbk sheet -mode ts signed_set_decomposition_existence.tex

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
	rm -f signed_set_decomposition_existence.*
	make reset
	bbk sheet -mode mk > Makefile