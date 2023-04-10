# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make measurable_function_operations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

measurable_function_operations.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > measurable_function_operations.tex

measurable_function_operations.pdf: ../../*.tex ../../trademark.pdf *.tex measurable_function_operations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode measurable_function_operations.tex
	../../latexrun measurable_function_operations.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: measurable_function_operations.pdf
	open measurable_function_operations.pdf

o:
	make open

terms: measurable_function_operations.tex
	bbk sheet -mode ts measurable_function_operations.tex

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
	rm -f measurable_function_operations.*
	make reset
	bbk sheet -mode mk > Makefile