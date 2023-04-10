# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make measure_vector_space.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

measure_vector_space.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > measure_vector_space.tex

measure_vector_space.pdf: ../../*.tex ../../trademark.pdf *.tex measure_vector_space.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode measure_vector_space.tex
	../../latexrun measure_vector_space.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: measure_vector_space.pdf
	open measure_vector_space.pdf

o:
	make open

terms: measure_vector_space.tex
	bbk sheet -mode ts measure_vector_space.tex

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
	rm -f measure_vector_space.*
	make reset
	bbk sheet -mode mk > Makefile