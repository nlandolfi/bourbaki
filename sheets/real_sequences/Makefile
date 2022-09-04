# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_sequences.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

real_sequences.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_sequences.tex

real_sequences.pdf: ../../*.tex ../../trademark.pdf *.tex real_sequences.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_sequences.tex
	../../latexrun real_sequences.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_sequences.pdf
	open real_sequences.pdf

o:
	make open

terms: real_sequences.tex
	bbk_sheet -mode ts real_sequences.tex

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
	rm -f real_sequences.*
	make reset
	bbk_sheet -mode mk > Makefile